
# swiff

Detects a SWF file's width/height given a url

### Usage

swiff is a command-line utility that accepts a single argument:

```bash
swiff https://example.org/example.swf
```

### Output

swiff always outputs JSON, the structure is as follows:


```javascript
{
  "success": true, /* or false */
  "errors": [], /* or a list of strings */
  "info": { /* or null */
    "width": 1280,
    "height": 720
  } /* except the actual width/height */
}
```

### Compatibility

swiff supports uncompressed SWF files, gzip-compressed SWF files, and lzma-compressed SWF files.

It only downloads as little as it needs to find the resolution, and doesn't store anything on disk.


