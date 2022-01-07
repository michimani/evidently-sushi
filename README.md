evidently-sushi
===

This is a sample for creating a CloudWatch Evidently project.

# Preparation

- Install AWS CLI

# Deploy

## via CloudFormation

```bash
sh ./cfn/10-evidently.sh deploy
```

# Start Launch

```bash
aws evidently start-launch \
--project FoodProject \
--launch SushiLaunch
```

# Check

## Project

```bash
aws evidently get-project \
--project FoodProject
```

## Feature

```bash
aws evidently get-feature \
--project FoodProject \
--feature SushiFeature
```

## Launch

```bash
aws evidently get-launch \
--project FoodProject \
--launch SushiLaunch
```

# Evaluate Feature

Run the "evaluate-feature" API 10 times with different EntityIDs.

```bash
for i in `seq 10`; do
  aws evidently evaluate-feature \
  --project FoodProject \
  --feature SushiFeature \
  --entity-id "$(date -u +'%s')-${i}" \
  --query 'value.stringValue' \
  --output text
  sleep 1
done
```

Then you will get the following output.

```bash
engawa
engawa
uni
tamago
maguro
maguro
tamago
uni
uni
maguro
```

If you specify the EntityID set in EntityOverrides, a fixed value will be returned regardless of the percentage of traffic.

```bash
for i in `seq 10`; do
  aws evidently evaluate-feature \
  --project FoodProject \
  --feature SushiFeature \
  --entity-id 'shari' \
  --query 'value.stringValue' \
  --output text
done
```

The output will look like the following

```bash
no neta
no neta
no neta
no neta
no neta
no neta
no neta
no neta
no neta
no neta
```
