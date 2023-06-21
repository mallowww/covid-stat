## Please create a simple go project that satisfy the following requirements

You're asked to write a simple JSON API to summarize COVID-19 stats using this public API, https://static.wongnai.com/devinterview/covid-cases.json

1. Your project must use Go, Go module, and Gin framework

2. You create a JSON API at this endpoint /covid/summary

3. The JSON API must count number of cases by provinces and age group

4. There are 3 age groups: 0-30, 31-60, and 60+ if the case has no age data, please count as "N/A" group

5. We encourage you to write tests, which we will give you some extra score

6. Please zip the whole project and upload to the form.