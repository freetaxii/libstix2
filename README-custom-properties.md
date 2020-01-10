# FreeTAXII/libstix2

## Handling Custom Properties

This note explains how to use custom properties with this library.


### Decoding Custom Properties

First, let’s deal with the decoding of an attack pattern JSON string that has a custom property called “some_custom_property”, where that custom property is of type string. 

You may have an attack pattern that looks like this:

```
{
    "type": "attack-pattern",
    "spec_version": "2.1",
    "id": "attack-pattern--d62e1eff-eb93-42e2-bd90-dabff3b93427",
    "created": "2018-06-05T18:25:15.917Z",
    "modified": "2018-06-05T18:25:15.917Z",
    "name": "Phishing",
    "aliases": ["Banking1", "ATM2"],
    "some_custom_property": "some_custom_value"
}
```

In your code you have consumed this JSON string and stored it in a []byte called “foo”. So you could decode it by doing:

```
o, _ := attackpattern.Decode(foo)
```

This will give you an attack pattern object called “o” and all of the custom properties will be stored in a property called “o.Custom” which is a `map[string]*json.RawMessage`. This means that we will have all of the key names for the custom properties, but all of their values will be untouched and left as a `[]byte`.  So we are not decoding the values at this stage. So if you do not know how to process them, you do not need to worry about it. Just store them as an array of bytes.

If you were to print out o.Custom you would get:

```
map[some_custom_property:[34 115 111 109 101 95 99 117 115 116 111 109 95 118 97 108 117 101 34]]
```

Then when you want to process that data in o.Custom, if you know how to, you would just unmarshal those properties, such as:

```
var data string
json.Unmarshal(o.Custom["some_custom_property"], &data)
```

If you were to print out value of “data” you would get:

```
some_custom_value
```


## Encoding Custom Properties

This will show you how too add custom properties to the objects that I have defined in this library. You would create a new type and embed the one from my library in it, like so:

```
type myCustomAttackPattern struct {
    *attackpattern.AttackPattern
    SomeCustomProperty string `json:"some_custom_propety,omitempty"`
}
```

Then in your main program you could do:

```
o := attackpattern.New()
customAP := myCustomAttackPattern{AttackPattern: o}
```

or


```
o := attackpattern.New()
var customAP myCustomAttackPattern
customAP.AttackPattern = o
```

At this point you can then populate your new custom property with data and encode it like so:

```
customAP.SomeCustomProperty = "some custom string data"
data, - := json.MarshalIndent(customAP, "", "    ")
```

Now if you print out “data” you will get an attack pattern with your new custom property. 

```
{
    "type": "attack-pattern",
    "spec_version": "2.1",
    "id": "attack-pattern--2b88c3ff-410d-44c4-9056-2ead668ab13c",
    "created": "2020-01-10T05:52:30.494Z",
    "modified": "2020-01-10T05:52:30.494Z",
    "some_custom_propety": "some custom string data"
}
```

## Working Examples

I have two examples in my examples directory that show this, 06-decode-custom-properties.go and 07-use-custom-properties.go



