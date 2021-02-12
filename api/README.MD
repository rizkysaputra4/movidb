## API ROUTE

#

### AUTH Endpoint

<br/>

    METHOD /auth

This route is used to handle auth related like check username, check email, register, etc. This route also open to public.

&nbsp;

#### Check Username

route to check if username exist. thiis route will giving error response when username exist

```
POST /auth/register-username
```

json body example:

```json
{
  "user_name": "movidb"
}
```

#### Check Email

route to check if email exist. this route will giving error response when email exist

```
POST /auth/register-email
```

json body example:

```json
{
  "email": "movidb@mail.com"
}
```

#### Register

route that added user into db.

```
POST /auth/register
```

json body example:

```json
{
  "user_name": "movidb",
  "email": "movidb@mail.com",
  "country_id": "id",
  "password": "movidb",
  "email": "movidb@mail.com"
}
```

#### Login / Check

to check if username or email exist.

```
POST /auth/login-username
```

json body example:

```json
{
  "user_name": "movidb"
}
```

or

```json
{
  "email": "movidb@mail.com"
}
```

#### Login / Check if Passowrd Match

to check if password match.

```
POST /auth/login-password
```

json body example:

```json
{
  "user_name": "movidb",
  "password": "movidb"
}
```

or

```json
{
  "email": "movidb",
  "password": "movidb"
}
```

#### Logout

Clean user cookies.

    GET /auth/logout

&nbsp;

## Admin Endpoint

This route only available for admin.

&nbsp;

#### Register new admin

Admin add new admin with registering new account.

```
POST /admin/register-new-admin
```

json body example:

```json
{
  "user_name": "newadmin",
  "email": "newadmin@mail.com",
  "country_id": "id",
  "password": "movidb",
  "email": "movidb@mail.com",
  "role": "6"
}
```

&nbsp;

#### CHANGE LEVEL

This route used for changing authority level (role).
note: admin cannot modify userID that have role higher than him, or change some role higher than his role.

```
PUT /admin/admin-level
```

json body example:

```json
{
  "user_id": 6,
  "role": 10
}
```

&nbsp;

#### New Identifier

Identifier used to identify whether review, comment, etc. is blocked, in review, online, etc.

    POST /admin/new-identifier

json body example:

```json
{
  "info": "blocked"
}
```

&nbsp;

#### Get Admin List

Get List of admin

    GET /admin/admin-list

&nbsp;

#### Search User

Search user

    GET /admin/user/search?uid=admin&limit=2&offset=0

&nbsp;

#### Add New Movie Type

Type is to identify is it movie, kdrama, or anime

    POST /admin/movie/add-new-movie-type

json body example:

```json
{
  "type_name": "tv-series"
}
```

&nbsp;

#### Add New Movie

Inserting new movie into db.

    POST /admin/movie/new-movie

json body example:

```json
{
  "movie_title": "iron man",
  "movie_synopsis": "synopis",
  "release_date": "10 jan 2010",
  "imdb_rating": 10,
  "imdb_num_vote": 1000,
  "metacritics": 4,
  "poster_link": "",
  "trailer_link": "youtube.com",
  "duration": "10 minutes",
  "awards": "",
  "type_id": 1,
  "country_id": "id",
  "language": "id",
  "unique_link": "iron-man",
  "identifier_id": 1
}
```

&nbsp;

#### Edit Movie

Edit Movie.

    PUT /admin/movie/edit-movie

json body example:

```json
{
  "movie_title": "iron man",
  "movie_synopsis": "synopis",
  "release_date": "10 jan 2010",
  "imdb_rating": 10,
  "imdb_num_vote": 1000,
  "metacritics": 4,
  "poster_link": "",
  "trailer_link": "youtube.com",
  "duration": "10 minutes",
  "awards": "",
  "type_id": 1,
  "country_id": "id",
  "language": "id",
  "unique_link": "iron-man",
  "identifier_id": 1
}
```

&nbsp;

#### Add New People Role

The table contains all role in movie industry like actress, actor, director, etc.

    POST /admin/movie/new-people-role

json body example:

```json
{
  "role_name": "writer"
}
```

&nbsp;

#### Update New People Role

The table contains all role in movie industry like actress, actor, director, etc.

    POST /admin/movie/update-people-role

json body example:

```json
{
  "role_id": 1,
  "role_name": "actress"
}
```

&nbsp;

### USER ENDPOINT

Route that accessible by user or registered member.

&nbsp;

#### Update user full info

Route to update user information like, birthdate, social media link, etc.

    PUT /user/{user-id}

```json
{
  "user_full_name": "john Doe",
  "birthdate": "1 January 2020",
  "bio": "movie lovers",
  "fb_link": "facebook.com",
  "twitter_link": "twitter.com",
  "ig_link": "ig.com",
  "sex": true
}
```

for sex variable: true equal to male, false equal to female

&nbsp;

### PUBLIC ENDPOINT

&nbsp;

#### Get My Role

GET /public/my-role

This will return your role in number 1 - 99
