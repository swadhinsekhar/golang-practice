package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
)

// ConfigWatcher holds a Viper instance and a channel for configuration changes
type ConfigWatcher struct {
	v             *viper.Viper
	configChange chan bool
}

// ReadConfig reads the configuration from the given file path
func ReadConfig(filePath string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(filePath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}

// NewConfigWatcher creates a new ConfigWatcher instance
func NewConfigWatcher(filePath string) (*ConfigWatcher, error) {

  v, err := ReadConfig(filePath)
	if err != nil {
		return nil, err
  }

	configChange := make(chan bool)
	watcher := &ConfigWatcher{
		v:             v,
		configChange:  configChange,
	}

	// Watch for changes in the config file
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		configChange <- true // Notify about config change
	})

	return watcher, nil
}

// UpdateConfig updates the specified keys in the provided Viper instance
func (c *ConfigWatcher) UpdateConfig(key string, value any) {
	c.v.Set(key, value)
}

// WriteConfig writes the updated configuration back to the file
func (c *ConfigWatcher) WriteConfig() error {
	return c.v.WriteConfig()
}

func (c *ConfigWatcher) GetString(key string) string {
	return c.v.GetString(fmt.Sprintf("%v", key))
}

func (c *ConfigWatcher) GetInt(key string) int {
	return c.v.GetInt(fmt.Sprint(key))
}

type HTTPApplication struct {
	Enable bool   `mapstructure:"enable"`
	Name   string `mapstructure:"name"`
	URLs   string `mapstructure:"urls"`
}

type HTTPSApplication struct {
	Enable bool   `mapstructure:"enable"`
	Name   string `mapstructure:"name"`
	URLs   string `mapstructure:"urls"`
}

type SSHApplication struct {
	Enable      bool   `mapstructure:"enable"`
	Name        string `mapstructure:"name"`
	Protocol    string `mapstructure:"protocol"`
}

func (c *ConfigWatcher)LoadHTTPApplication() ([]HTTPApplication, error) {
	var httpApplications []HTTPApplication
	if err := c.v.UnmarshalKey("applications.http", &httpApplications); err != nil {
		fmt.Println("Error reading HTTP applications:", err)
    return nil, fmt.Errorf("error reading config file: %w", err)
	}
  return httpApplications, nil
}
func (c *ConfigWatcher) LoadHTTPSApplication() ([]HTTPSApplication, error) {
	var httpsApplications []HTTPSApplication
	if err := c.v.UnmarshalKey("applications.https", &httpsApplications); err != nil {
		fmt.Println("Error reading HTTP applications:", err)
    return nil, fmt.Errorf("error reading config file: %w", err)
	}
  return httpsApplications, nil
}

func (c *ConfigWatcher) LoadSSHApplication() ([]SSHApplication, error) {
	var sshApplications []SSHApplication
	if err := c.v.UnmarshalKey("applications.ssh", &sshApplications); err != nil {
		fmt.Println("Error reading HTTP applications:", err)
    return nil, fmt.Errorf("error reading config file: %w", err)
	}
  return sshApplications, nil
}

func main() {
	filePath := "config.json"

  configWatcher, err := NewConfigWatcher(filePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

  logrus.Println(configWatcher.GetString("webpathserver.address"))
  logrus.Println(configWatcher.GetInt("webpathserver.port"))
  logrus.Println(configWatcher.GetString("node.node_id"))
  logrus.Println(configWatcher.GetString("node.tunnel_id"))
  /*
	configWatcher.UpdateConfig("key1", "new_value1")
	configWatcher.UpdateConfig("key2", 456)
	configWatcher.UpdateConfig("key3", 357)
  */
	// Write the updated configuration back to the file
	if err := configWatcher.WriteConfig(); err != nil {
		fmt.Println("Error writing config file:", err)
		return
	}

  httpApplications, _ := configWatcher.LoadHTTPApplication()
	for _, app := range httpApplications {
    fmt.Printf("Application: %s, Enable: %t, URLs: %s\n", app.Name, app.Enable, app.URLs)
  }

  httpsApplications, _ := configWatcher.LoadHTTPSApplication()
	for _, app := range httpsApplications {
    fmt.Printf("Application: %s, Enable: %t, URLs: %s\n", app.Name, app.Enable, app.URLs)
  }

  sshApplications, _ := configWatcher.LoadSSHApplication()
	for _, app := range sshApplications {
    fmt.Printf("Application: %s, Enable: %t, URLs: %s\n", app.Name, app.Enable, app.Protocol)
  }

  /*
  // Run the main loop to listen for config changes
	for {
		select {
		case <-configWatcher.configChange:
			fmt.Println("Handling config change...")
			// Handle configuration update here
			// For example:
			// newConfig := configWatcher.v.GetString("key")

		// Perform other tasks if needed
		}
	}
  */
}

