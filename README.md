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

