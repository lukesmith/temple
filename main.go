package main

import (
  "os"
  "io/ioutil"
  "html/template"
	"log"
  "github.com/urfave/cli"
  "gopkg.in/yaml.v2"
)

func main() {
  app := cli.NewApp()
  app.Name = "Temple"
  app.Usage = "Template files"
	app.Version = "0.0.1"

  app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "template",
			Usage: "The path to the file to template.",
		},
		cli.StringSliceFlag{
			Name: "data",
			Usage: "The path to a yaml file containing the data to use to template.",
		},
	}

  app.Action = func(c *cli.Context) error {
  	templ, err := ioutil.ReadFile(c.String("template"))
  	if err != nil {
  		return err
  	}

		dataFiles := c.StringSlice("data")
		data := make(map[interface{}]interface{})

		for _, envFile := range dataFiles {
			s, err := ioutil.ReadFile(envFile)
			if err != nil {
				return err
			}

			d := make(map[interface{}]interface{})

			err = yaml.Unmarshal([]byte(s), &d)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			for k, v := range d {
				data[k] = v
			}
		}

    t := template.New("Template")
    t, err = t.Parse(string(templ))
    if err == nil {
    	t.Execute(os.Stdout, data)
    } else {
    	log.Fatalf("error: %v", err)
		}

    return nil
  }

  app.Run(os.Args)
}
