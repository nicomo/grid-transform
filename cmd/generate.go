/*
Copyright © 2019 NICOLAS MORIN <nicolas.morin@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate a dgraph compatible json file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate called")
		genDgraphJSON()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func genDgraphJSON() {

	// 1. read json source file
	file, err := os.Open("json-source-example.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//TODO: create target JSON File

	// decode source file
	dec := json.NewDecoder(file)

	// read & discard open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {

		// decode json into struct
		instSrc := InstituteSource{}
		err := dec.Decode(&instSrc)
		if err != nil {
			log.Fatal(err)
		}

		//transform into target JSON
		instT := transform(instSrc)
		fmt.Printf("%+v\n", instT)

		// add to target JSON file

	}
}

func transform(instSrc InstituteSource) InstituteTarget {

	// create target struct, init with flat values
	instT := InstituteTarget{
		UID:          "_:" + instSrc.ID,
		DgraphType:   "institute",
		Name:         instSrc.Name,
		WikipediaURL: instSrc.WikipediaURL,
		Status:       instSrc.Status,
		Established:  instSrc.Established,
		Acronyms:     instSrc.Acronyms,
	}

	// we take the 1st link if it exists
	if len(instSrc.Links) > 0 {
		instT.Link = instSrc.Links[0]
	}

	// build the address / location from 1st address
	if len(instSrc.Addresses) > 0 {
		instT.Location = Location{
			Type:        "Point",
			Coordinates: []float64{instSrc.Addresses[0].Lat, instSrc.Addresses[0].Lng},
		}
		instT.City = instSrc.Addresses[0].City
		instT.Country = instSrc.Addresses[0].Country
		instT.CountryCode = instSrc.Addresses[0].CountryCode
		instT.GeonamesCity = instSrc.Addresses[0].GeonamesCity.GCID
	}

	// Parents and Children
	p := []Parent{}
	c := []Child{}
	for _, v := range instSrc.Relationships {
		switch v.Type {
		case "Parent":
			p = append(p, Parent{UID: v.ID})
		case "Child":
			c = append(c, Child{UID: v.ID})
		}
	}

	// external ids, starting with grid_id as an external ID
	instT.Xids = []Xids{
		Xids{
			Source: "grid",
			Xid:    instSrc.ID,
		},
	}

	// parsing other external IDs
	// this is ugly but so is the source data
	for k, v := range instSrc.ExternalIds {
		for kk, vv := range v.(map[string]interface{}) {
			if kk == "all" && vv != nil && reflect.TypeOf(vv).Kind() == reflect.Slice {
				s := reflect.ValueOf(vv)
				instT.Xids = append(instT.Xids,
					Xids{
						Source: k,
						Xid:    fmt.Sprint(s.Index(0)),
					},
				)
			}
		}
	}

	return instT
}
