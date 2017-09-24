# Nettaton
React frontend in progress, until then you can use curl like so:
```sh
$ curl -s https://nettaton.com/question
```
This should get you a question that looks like a minified version of this:
```sh
{
    "ipAddress":"10.10.10.10",
    "network":"24",
    "questionKind":"last"
}
```
This is asking for the last valid IP address in the network that 10.10.10.10/24 is in. Sometimes the question have prefix notation, sometimes subnet mask.
The different types of question are:

* First - first valid IP of network
* Last - last valid IP of network
* Broadcast - broadcast IP of network
* FirstAndLast - first and last valid IP addresses of network (e.g. 10.10.10.1-10.10.10.255)
* HostsInNet - how many valid hosts there are in the network

To answer, just post your "answer" to the endpoint, like so:
```sh
curl -s -d '{"ipAddress":"10.10.10.10","network":"24","questionKind":"first","answer":"10.10.10.1"}' \
  -H "Content-Type: application/json" \
  -X POST https://nettaton.com/question | json_pp
```
The response should look like this:
```sh
{
    "userAnswer" : "10.10.10.1",
    "actualAnswer" : "10.10.10.1",
    "marks" : null
}
```
If your userAnswer is the same as the actualAnswer, you got it right!