$.ajax({
    type: "POST",
    url: "https://cp.pushwoosh.com/json/1.3/createMessage",
    data: JSON.stringify({
        "request": {
            "application": "APPLICATION CODE",
            "auth": "API TOKEN",
            "notifications": [{
                "send_date": "now",
                "ignore_user_timezone": true,
                "content": "Hello world!"
            }]
        }
    }),
    dataType: "json"
}).done(function(data) {
    console.log(data);
});