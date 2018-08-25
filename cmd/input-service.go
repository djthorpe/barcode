package main

import (
	"os"

	// Frameworks
	"github.com/djthorpe/gopi"

	// Modules
	_ "github.com/djthorpe/gopi-input/sys/input"
	_ "github.com/djthorpe/gopi-input/sys/keymap"
	_ "github.com/djthorpe/gopi/sys/logger"
	_ "github.com/djthorpe/gopi/sys/rpc/grpc"
	_ "github.com/djthorpe/gopi/sys/rpc/mdns"

	// RPC Services
	_ "github.com/djthorpe/gopi-input/rpc/grpc/input"
)

/*
var (
	start = make(chan struct{})
)
*/

/*
func PrintDevicesTable(devices []gopi.InputDevice) {
	// Table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Name", "Bus"})
	for _, d := range devices {
		table.Append([]string{
			fmt.Sprint(d.Type()),
			d.Name(),
			fmt.Sprint(d.Bus()),
		})
	}
	table.Render()
}

func EventLoop(app *gopi.AppInstance, done <-chan struct{}) error {
	// Subscribe to events
	evt_input := app.Input.Subscribe()

FOR_LOOP:
	for {
		select {
		case <-start:
			app.Logger.Info("Start")
		case <-done:
			app.Logger.Info("Done")
			break FOR_LOOP
		case event := <-evt_input:
			app.Logger.Info("Input: %v", event)
		}
	}

	// Unsubscribe from events
	app.Input.Unsubscribe(evt_input)

	// Return success
	return nil
}

func Main(app *gopi.AppInstance, done chan<- struct{}) error {
	// Open devices
	if devices, err := app.Input.OpenDevicesByName("", gopi.INPUT_TYPE_ANY, gopi.INPUT_BUS_ANY); err != nil {
		return err
	} else {
		PrintDevicesTable(devices)
	}

	if watch, _ := app.AppFlags.GetBool("watch"); watch {
		// Send start flag
		start <- gopi.DONE

		// Wait for CTRL+C
		fmt.Println("Watching for events, press CTRL+C to end")
		app.WaitForSignal()
		fmt.Println("Terminating")
	}

	done <- gopi.DONE
	return nil
}
*/

func main() {
	// Create the configuration
	config := gopi.NewAppConfig("rpc/service/input:grpc")

	// Set the RPCServiceRecord for server discovery
	config.Service = "input"

	// Run the server and register all the services
	os.Exit(gopi.RPCServerTool(config))
}

/*
func main() {
	config := gopi.NewAppConfig("input")
	config.AppFlags.FlagBool("watch", false, "Watch for device events")
	os.Exit(gopi.CommandLineTool(config, Main, EventLoop))
}
*/
