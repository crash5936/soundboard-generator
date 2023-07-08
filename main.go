package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/BurntSushi/toml"
)

type Sound struct {
	Name, Description string
}

type SoundsToml struct {
	Sounds []Sound
}

func main() {
	prepare_sounds_config("sounds/")
	// generate_html()
}

func generate_html() {

	sound_names := get_sound_names("sounds/")

	var sound_structs []Sound

	for _, name := range sound_names {
		sound_structs = append(sound_structs, Sound{name, name})
	}

	tmpl_str, err := os.ReadFile("templates/soundboard.tmpl")

	tmpl, err := template.New("soundboard").Parse(string(tmpl_str))

	if tmpl == nil {
		panic("Failed parsing tmpl.")
	}

	check_error(err)

	f, err := os.Create("output/output.html")
	
	check_error(err)
	
	err = tmpl.Execute(f, sound_structs)

	check_error(err)

	err = f.Close()

	check_error(err)
}

func prepare_sounds_config(sounds_path string) {
	sound_names := get_sound_names(sounds_path)

	fmt.Println(sound_names)

	var sounds []Sound
	var config SoundsToml

	for _, name := range sound_names {
		sounds = append(sounds, Sound{name, name})
	}

	config = SoundsToml{Sounds: sounds}

	f, err := os.Create(filepath.Join(sounds_path, "sounds.toml"))
	check_error(err)

	err = toml.NewEncoder(f).Encode(config)
	check_error(err)

	err = f.Close()
	check_error(err)
}

func get_sound_names(file_path string) []string {
	files, err := os.ReadDir(file_path)
	check_error(err)
	re := regexp.MustCompile(`^(.*)\.mp3$`)
	var sound_names []string
	for _, sound := range files {
		file_name := sound.Name()
		sound_submatch := re.FindStringSubmatch(file_name)
		if len(sound_submatch) == 2 {
			sound_name := sound_submatch[1]
			sound_names = append(sound_names, sound_name)
		}
	}
	return sound_names
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}