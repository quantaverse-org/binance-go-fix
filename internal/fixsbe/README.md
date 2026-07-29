# FIX SBE generated codecs

The Go files in this directory are generated from:

```text
schemas/spot-fixsbe-1_1.xml
```

Generator: SbeTool 1.37.1.

```bash
java --add-opens java.base/jdk.internal.misc=ALL-UNNAMED \
  -Dsbe.target.language=Golang \
  -Dsbe.target.namespace=fixsbe \
  -Dsbe.output.dir=internal \
  -Dsbe.validation.stop.on.error=true \
  -jar sbe-all-1.37.1.jar \
  schemas/spot-fixsbe-1_1.xml
```

SbeTool reports a warning for Binance's `uint32` `MDEntries.numInGroup`;
the schema intentionally uses `groupSize32Encoding` for trade batches.

SbeTool 1.37.1 also shadows the `InstrumentList` receiver in three generated
loops. After regeneration, rename those loop indexes from `i` to `idx` before
committing the generated package.
