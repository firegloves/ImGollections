# ImGollection

ImGollection is a tool to automate the creation os series of images starting by collections of layers.

By providing a list of folders identifying different layers, ImGollection will produce all the combinations resulting from overlapping the various layers.

In the `resources/config.yml` file, the `layerfolders` property lists the folders representing the various layers.
ImGollection will process them in the order specified from top to bottom.

## Example Flow

With this configuration

```
layerfolders:
- "input/Default/Skin"
- "input/Default/Eyes"
- "input/Default/Mouth"
```

ImGollection will execute the current stream:

For each file present in the `input/Default/Skin` folder.
    For each file present in the `input/Default/Eyes` folder
        For each file present in the `input/Default/Mouth` folder
            Overlay the layers to obtain a composite image
            Save the resulting image

The target folder where the images will be generated is specified in the `config.yaml`'s `outputfolder` folder property