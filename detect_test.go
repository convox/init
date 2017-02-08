package init

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func v(dir string) string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s/%s:/tmp/app", wd, dir)
}

func TestRubyDetect(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("ruby-getting-started"), "convox/init", "detect").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, string(out), "Ruby\n")
}

func TestRubyRelease(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("ruby-getting-started"), "convox/init", "release").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, `---
addons:
- heroku-postgresql
config_vars:
  LANG: en_US.UTF-8
  RAILS_ENV: production
  RACK_ENV: production
  SECRET_KEY_BASE: 9e21b5bb3ed7c53fb2c3bcfe1f06712a4ea39b487a8e0965551af8c4e2155637a59dcda5e06937fe351e810b1fbe57b83a353085f591ad893632f3a6e200f4b4
  RAILS_SERVE_STATIC_FILES: enabled
default_process_types:
  rake: bundle exec rake
  console: bin/rails console
  web: bin/rails server -p $PORT -e $RAILS_ENV
  worker: bundle exec rake jobs:work
`,
		string(out))
}

func TestNodeDetect(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("node-js-getting-started"), "convox/init", "detect").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, string(out), "Node.js\n")
}

func TestNodeRelease(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("node-js-getting-started"), "convox/init", "release").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, `addons: []
default_process_types:
  web: npm start
`,
		string(out))
}

func TestPythonDetect(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("python-getting-started"), "convox/init", "detect").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, string(out), "Python\n")
}

func TestPythonRelease(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("python-getting-started"), "convox/init", "release").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, `---
config_vars:


addons:
  - heroku-postgresql
`,
		string(out))
}

func TestJavaDetect(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("java-getting-started"), "convox/init", "detect").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, string(out), "Java\n")
}

func TestJavaRelease(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("java-getting-started"), "convox/init", "release").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, `---
addons:
  - heroku-postgresql
`,
		string(out))
}

func TestPHPDetect(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("php-getting-started"), "convox/init", "detect").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, string(out), "PHP\n")
}

func TestPHPRelease(t *testing.T) {
	out, err := exec.Command("docker", "run", "-v", v("php-getting-started"), "convox/init", "release").CombinedOutput()
	assert.NoError(t, err)
	assert.Equal(t, `---
addons: []
default_process_types:
  web: php -S 0.0.0.0:$PORT
`,
		string(out))
}
