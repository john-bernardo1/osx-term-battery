osx-term-battery
================

<table>
  <tr>
    <td>Binary</td>
    <td>osx_battery</td>
  </tr>
  <tr>
    <td>Description</td>
    <td>Displays visual battery information in your OSX zsh terminal</td>
  </tr>
  <tr>
    <td>Go Version</td>
    <td>>= 1.2</td>
  </tr>
</table>

# Install

Make sure you have [Google Go >= 1.2](http://golang.org/doc/install) installed. May work with older
versions but has not been tested.

```
cd osx-term-battery
go install
```

Then, open up your desired .zsh-theme file in a text editor and change or add this line:
```
function battery_charge {
  echo `$GOPATH/bin/osx_battery`
}
RPROMPT='$(battery_charge)'
```

# Info

This is merely a port of [Steve Losh's batcharge.py](http://stevelosh.com/blog/2010/02/my-extravagant-zsh-prompt/)
