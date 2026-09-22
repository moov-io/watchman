#!/usr/bin/env python3
"""Restore the Watchman splash after fyne package regenerates cmd/ui/wasm/index.html."""

import pathlib
import re
import sys


# Fyne's wasm backend (fyne-io/glfw-js) listens on document keydown/keyup and
# calls preventDefault for every key. That cancels the browser shortcuts Chrome
# still lets a page override, including Cmd/Ctrl+1-9 and reload. Ctrl+Tab is
# reserved, so it keeps working. This capture-phase listener runs first and
# keeps those shortcuts from reaching that handler.
# Refs fyne-io/fyne#6541
BROWSER_KEYS = """\
\t\t<script id="watchman-browser-keys">
\t\t\t(function () {
\t\t\t\tfunction browserShortcut(e) {
\t\t\t\t\tvar key = e.key || "";
\t\t\t\t\tif (key === "F5") {
\t\t\t\t\t\treturn true;
\t\t\t\t\t}
\t\t\t\t\tif (!(e.metaKey || e.ctrlKey)) {
\t\t\t\t\t\treturn false;
\t\t\t\t\t}
\t\t\t\t\tif (key.length === 1 && key >= "1" && key <= "9") {
\t\t\t\t\t\treturn true;
\t\t\t\t\t}
\t\t\t\t\treturn key === "r" || key === "R";
\t\t\t\t}
\t\t\t\tfunction release(e) {
\t\t\t\t\tif (browserShortcut(e)) {
\t\t\t\t\t\te.stopImmediatePropagation();
\t\t\t\t\t}
\t\t\t\t}
\t\t\t\twindow.addEventListener("keydown", release, true);
\t\t\t\twindow.addEventListener("keyup", release, true);
\t\t\t})();
\t\t</script>
"""


def main() -> None:
    html_path = pathlib.Path(sys.argv[1])
    html = html_path.read_text()

    html = re.sub(
        r'(<div class="application-name">).*?(</div>)',
        r"\1Watchman\2",
        html,
        count=1,
        flags=re.DOTALL,
    )
    if "<title>" not in html:
        html = html.replace("<head>", "<head>\n\t\t<title>Watchman</title>", 1)

    if 'id="watchman-browser-keys"' not in html:
        html = html.replace("<head>", "<head>\n" + BROWSER_KEYS, 1)

    if "var wasmResponse = fetch(" not in html:
        match = re.search(r'instantiateStreaming\(fetch\("([^"]+\.wasm)"\)', html)
        if match:
            wasm = match.group(1)
            html = html.replace(
                f'instantiateStreaming(fetch("{wasm}")',
                "instantiateStreaming(wasmResponse",
                1,
            )
            needle = 'action.innerHTML = "Downloading wasm_exec.js"'
            html = html.replace(
                needle,
                f'var wasmResponse = fetch("{wasm}");\n\t\t\t\t\t\t{needle}',
                1,
            )

    html_path.write_text(html)


if __name__ == "__main__":
    main()
