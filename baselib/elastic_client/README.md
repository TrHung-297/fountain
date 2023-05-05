# Config for ELASTIC

## EXAMPLE

```json
{
    "Elastic": {
        "Name": "string",
        "Environment": "string", // if not set, set by Env.Environment
        "Hosts": "string", // can be a array with split by comma or semi-colon; example.com;http://es.example.com
        "Username": "string",
        "Password": "string",
        "AppKey": "string",
        "SecretKey": "string",
        "PlatformDefault": "string", // if not set, set by Env.PlatformDefault; example: GTV_Plus
        "LogTypeDefault": "string" // if not set, set by Env.LogTypeDefault; example: LogEvent
    }
}
```