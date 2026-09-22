#!/usr/bin/env python3
"""Restore the Watchman splash after fyne package regenerates cmd/ui/wasm/index.html."""

import pathlib
import re
import sys


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
