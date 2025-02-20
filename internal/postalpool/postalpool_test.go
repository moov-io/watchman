package postalpool

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/watchman/pkg/address"

	"github.com/stretchr/testify/require"
)

func TestPostalPool(t *testing.T) {
	svc := setupPostalPool(t)

	addr, err := svc.ParseAddress(context.Background(), "123 First St Anytown, CA 90210")
	require.NoError(t, err)
	require.NotEmpty(t, addr.Line1)

	t.Logf("%#v", addr)
}

// Benchmark_PostalPool tests the performance / memory usage of HTTP calls to a postal-server binary
// ran under the postalpool worker group.
//
// In order to accurately measure postal-server performance with libpostal you must build the binary with
//
//	GOTAGS="-tags libpostal" make build
//
// Otherwise your test will be measuring performance without libpostal compiled into postal-server.
func Benchmark_PostalPool(b *testing.B) {
	ctx := context.Background()
	svc := setupPostalPool(b)

	inputs := []string{
		"Flat 7B, Tower 2, Ocean Financial Centre, 12 Marina Boulevard, Singapore 018982",
		"Room 1403, West Wing, Trading Complex No. 5, 47 Al Souq Street, Dubai, United Arab Emirates",
		"Office 892, Floor 8, Edificio Comercial Torres, Avenida Balboa y Calle 42, Panama City, Panama",
		"Unit 15, 3rd Floor, 123 Pyongyang Industrial Zone, Rangnang District, Pyongyang, DPRK",
		"Suite 405, Business Center Red Square, 17 Tverskaya Street, Moscow 125009, Russian Federation",
		"Warehouse 23, Port Zone B, Terminal 4, Latakia Port Complex, Latakia, Syria",
		"Office 78, Tehran Trade Tower, Block 2, Valiasr Street, Tehran 19395-4791, Iran",
		"Villa 15, Street 7, Block 4, Diplomatic Quarter, Caracas 1010, Venezuela",
		"Room 2201, Finance Plaza Building, 333 Lujiazui Ring Road, Shanghai 200120, China",
		"Suite 17, Victoria Business Park, 45 Harare Drive, Harare, Zimbabwe",
		"Office Complex Delta, Building C, Floor 5, 89 Minsk Boulevard, Minsk 220114, Belarus",
		"Unit 908, Golden Trade Center, 78 Yangon Port Road, Yangon 11181, Myanmar",
		"Floor 3, Al-Zawra Tower, Block 215, Baghdad Commercial District, Baghdad, Iraq",
		"Building 45, Industrial Zone 3, Damascus International Airport Road, Damascus, Syria",
		"Suite 301, Havana Trade Building, 67 Malecon Avenue, Havana 10400, Cuba",
		"Office 12, Floor 4, Conakry Commerce Center, Route du Niger, Conakry, Guinea",
		"Unit 55, Khartoum Business Complex, Al Gamhoria Avenue, Khartoum, Sudan",
		"Room 789, Floor 7, Trade Tower 3, Kim Il Sung Square, Pyongyang, DPRK",
		"Building 23, Floor 2, Sevastopol Maritime Complex, 45 Port Street, Sevastopol 99011",
		"Office 445, Tripoli Trade Center, Omar Al-Mukhtar Street, Tripoli, Libya",
		"c/o Purepecha Trucking Co. Uruapan, Michoacan Mexico",
		"Calle 7 Sur No. 42-70 of. 1105 Medellin  Colombia",
		"537/1-Z Defense Housing Area (DHA) Lahore  Pakistan",
		"Calle 10A 11A-02 Maicao  Colombia",
		"Calle Lago de La Doga 5312 Tijuana, Baja California Mexico",
		"Col Bella Vista, Casa No. 09, Camino a Rotulo de Coca Cola San Pedro Sula, Cortes Honduras",
		"Mexialtzingo 1964, Col. Americana Guadalajara, Jalisco 44150 Mexico",
		"Miguel Hidalgo Norte, No. 212, Zona Centro Allende, Coahuila Mexico",
		"Calle Enrique Cavazos No. 2326, Colonia Universidad Saltillo, Coahuila Mexico",
		"Number 3, 12 Narenjestan Street, Pasdaran Avenue Tehran  Iran",
		"No. 110 Baiyun Street Dalian, Liaoning China",
		"Blvd. Guillermo Batiz Paredes No. 1100, Col. Buenos Aires Culiacan, Sinaloa C.P. 80199 Mexico",
		"Ul. Bolshaya Tatarskaya D. 35, Str. 5 Moscow 115184 Russia",
		"Room 502, 90, Ponghak-dong, Pyongchon-guyok Pyongyang  Korea, North",
		"26-14 Hanakuma-cho Chuo-ku Kobe-shi, Hyogo-ken Japan",
		"Al-Mukalla Branch, Al-Kabas, Near Al-Mukalla Post Office Al-Mukalla, Hadhramout Yemen",
		"d. 25 ofis 13, 14, per. Avtomobilny Rostov-on-Don, Rostovskaya Oblast 344038 Russia",
	}
	b.ResetTimer()

	b.Run(fmt.Sprintf("Both/%s", svc.Ratio()), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			var index atomic.Int32
			for pb.Next() {
				// Get next address in a thread-safe way
				input := inputs[int(index.Add(1))%len(inputs)]
				addr, err := svc.ParseAddress(ctx, input)
				if addr.Line1 == "" {
					b.Fatalf("empty line1: %v", input)
				}
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("Client", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			var index atomic.Int32
			for pb.Next() {
				// Get next address in a thread-safe way
				input := inputs[int(index.Add(1))%len(inputs)]
				addr, err := svc.client.parseAddress(ctx, input, false)
				if addr.Line1 == "" {
					b.Fatalf("empty line1: %v", input)
				}
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	})

	b.Run("Direct", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			var index atomic.Int32
			for pb.Next() {
				// Get next address in a thread-safe way
				input := inputs[int(index.Add(1))%len(inputs)]
				addr := address.ParseAddress(ctx, input)
				if addr.Line1 == "" {
					b.Fatalf("empty line1: %v", input)
				}
			}
		})
	})
}

func setupPostalPool(tb testing.TB) *Service {
	tb.Helper()

	logger := log.NewTestLogger()

	conf := Config{
		Enabled:          true,
		Instances:        1,
		StartingPort:     10000,
		StartupTimeout:   60 * time.Second,
		BinaryPath:       filepath.Join("..", "..", "bin", "postal-server"),
		RequestTimeout:   5 * time.Second,
		CGOSelfInstances: 2,
	}
	_, err := os.Stat(conf.BinaryPath)
	if err != nil {
		if os.IsNotExist(err) {
			tb.Skipf("%s is not written, try `make build` before running this test again", conf.BinaryPath)
		}
	}

	svc, err := NewService(logger, conf)
	require.NoError(tb, err)

	tb.Cleanup(func() {
		svc.Shutdown()
	})

	return svc
}
