

### Prerequisite

1. Download the necessary configure files from aws: aws s3 cp s3://de012a78750e59b808d922b39535e862/v2.0.0-RC4-fork.5.tgz v2.0.0-RC4-fork.5.tgz

2. untar v2.0.0-RC4-fork.5.tgz and copy `config` dir into `full_prover_config`

### start

1. First start zkevm-prover-db: `docker-compose up -d zkevm-prover-db`;

2. Then start zkevm-prover: `docker-compose up -d zkevm-prover`

### benchmark

#### keccak256 test

- change test.node.config.toml
    - `TimestampResolution`: this field determines the timeout of a batch, change to a bigger number such as 100s
    - change `MTClient` and `Executor` URI field

- cd test; make run

- cd scripts/deploy_hash_contract; go run main.go


#### erc20 tranfer benchmark

- annotate `$(RUNZKPROVER)`

- modify as following:

```
diff --git a/test/benchmarks/sequencer/common/setup/setup.go b/test/benchmarks/sequencer/common/setup/setup.go
index 7eec5fde..f64b74a1 100644
--- a/test/benchmarks/sequencer/common/setup/setup.go
+++ b/test/benchmarks/sequencer/common/setup/setup.go
@@ -135,8 +135,21 @@ func Components(opsman *operations.Manager) error {
        if err != nil {
                return err
        }
-       time.Sleep(sleepDuration)
+       err = operations.StartComponent("agg")
+       if err != nil {
+               return err
+       }
+
+       err = operations.StartComponent("eth-tx-manager")
+        if err != nil {
+                return err
+        }

+       err = operations.StartComponent("sequence-sender")
+       if err != nil {
+               return err
+       }
+       time.Sleep(sleepDuration)
        return nil
 }
```