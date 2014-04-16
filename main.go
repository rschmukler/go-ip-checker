package main

import (
  "os"
  "log"
  "github.com/codegangsta/cli"
  "./app"
  "./providers"
)

const VERSION = "0.1.0"


func init() {
  log.SetFlags(0)
}

func main() {
  app := app.NewDDNSApp()
  app.Version = VERSION
  app.Commands = []cli.Command {
    {
      Name: "update",
      Usage: "update <domain>",
      Action: update(app),
      Flags: []cli.Flag {
        cli.StringFlag{Name: "ip-address, i", Usage: "IP Address to report"},
        cli.StringFlag{"provider, p", "iwantmyname.com", "API provider"},
        cli.BoolFlag{"configure, c", "Reconfigure"},
      },
    },
  }
  app.Run(os.Args)

}

func update(app *app.DDNSApp) func(c *cli.Context) {
  return func(c *cli.Context) {
    reconfigure := c.Bool("configure")

    provider, present := providers.GetProvider(c.String("provider"))

    if !present {
      log.Fatalf("Could not find provider specified (%s)\nValid Providers are:\n%s", c.String("provider"), providers.ListProviders())
    }

    config, present := app.Config[c.String("provider")]

    if !present || reconfigure {
      log.Printf("Please enter the following information: ")
      provider.GenerateConfig(app.Config)
      app.SaveConfig()
      config = app.Config[c.String("provider")]
    }

    ip := c.String("ip")

    provider.ReadConfig(config)
    go provider.Update(ip, app.Updates)
    printUpdate(<- app.Updates)
  }
}

func printUpdate(update app.DDNSUpdates) {
  log.Printf("[%s] (%s) %s\n", update.Type, update.From, update.Message)
}
