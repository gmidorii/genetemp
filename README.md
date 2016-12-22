## README

### Overview
genetemp generates [Java] file.

##### config file
- param setting
- yaml

##### template file
- template
- param is enclosed in bracket[]
```
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