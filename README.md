## README

### Overview
genetemp generates [Java] file.

##### config file
- param setting
- yaml
```
name(Required): GeneTemp
path(Required): github.com/midorigreen/genetemp
```

##### template file
- template
- param is enclosed in bracket[]
```
package [path]

public class [Name] {

}
```

### Build
```
go build
```

### Run
```
./genetemp -t [templatefile] -c [configfile]
```