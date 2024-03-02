# Allure Results local stats

## Motivation

- get local statistics from [allure2 results](https://allurereport.org/docs/how-it-works-test-result-file/) without building of the report
- stats can be used to send notifications

## Caveats

Current implementation is very primitive to exercise go programming language, and it only shows the stats for tests without considering retries (historyId is ignored).

This means, if there are several executions of the same test, then the status will be counted several times.

## Usage

- Declare `ALLURE_RESULTS` env variable, to show where allure results will be stored
- Download latest release to your pipeline and make it executable

```shell
    - wget https://github.com/cheshi-mantu/go-allure-result-stats/releases/latest/download/allurestat_linux_amd64 -O ./allurestat
    - chmod +x ./allurestat
```

- After tests execution add `allurestat` to get the PASS/FAIL stats.

```shell
$(./allurestat)
```

- Application execution of the app as shown above will result in creation of several environment variables with prefix `ALLURE_STAT_XXX`, where `XXX` could be as follows
  - PASSED
  - FAILED
  - BROKEN
  - SKIPPED

#### Example

Gitlab pipeline prints the stats for 8 successfully passed tests.

```shell
$ printenv | grep ALLURE_STAT_
ALLURE_STAT_PASSED=8
```