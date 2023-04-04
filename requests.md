# Requests and responces examples


### Delete album
Responce
```json
{
	"deleted": true
}
```

***In case of error***
```json
{
	"deleted": false,
	"error": "album not found"
}
```

### Add album
Request
```json
{
	"title": "Graduation",
	"author": "Kanye West",
	"year": 2007,
	"description": "Graduation is the third studio album from Chicago rapper Kanye West. The album was inspired by stadium tours, indie rock and house-music and was a huge departure from the sound Ye had used on his first two studio albums The College Dropout and Late Registration which featured samples and inspirations from soul and orchestral music. This album, on the other hand, included a much more electronic sound, featuring the use of layering synthesizers. Lyrically, Kanye analyzes himself and talks about his life after becoming famous and how he is criticized from the media.\nThe album famously faced off against 50 Cent’s Curtis in a sales war which Kanye won with 957,000 copies sold in the first week.\nSome critics argue this is the day gangsta rap was finally defeated.",
	"duration": "00:54:00",
	"songs": [
		{
			"title": "Good Morning",
			"duration": "3:15"
		},
		{
			"title":"Champion",
			"duration": "2:48"
		}
	],
	"geniousLink": "www.genius.com/albums/Kanye-west/Graduation"
}
```

Responce
```json
{
	"added": true,
	"id": "642b2b49120b5ffc18a4d0b2"
}
```

#### In case of error
```json
{
	"added" : false,
	"error": "unable to add album Graduation, this album already exists"
}
```

### Add admin

Request
```json
{
	"id": "6429a785208b8ed7434433b114",
}
```

Responce
```json
{
	"added": true,
}
```

***In case of error***
```json
{
	"added": false,
	"error": "the provided hex string is not a valid ObjectID"
}
```

### Register
Request
```json
{
	"username": "admin",
	"password": "admin"
}
```

Responce
```json
{
	"id": "642c282f9e5d01aed4246941"
}
```

***In case of error***
```json
{
	"error": "this user already exists"
}
```

### Login
Request
```json
{
	"username": "admin",
	"password": "admin"
}
```

Responce
```json
{
	"token": "_returns_token_"
}
```

***In case of error***
```json
{
	"error": "wrong username or password"
}
```
