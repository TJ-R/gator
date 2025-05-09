# gator

# Dependencies
 - Go
 - Postgres

# How to install
You must have both go and postgres installed.

After this you can clone the repository wherever on your device
in the clone repository run go install gator to download all of the necessary dependencies
and install the application

# Setup
Create a config file in home directory called .gatorconfig.json with content like
`
    {
        "db_url":"postgres://<username>:<password>@localhost:5432/gator?sslmode=disable",
        "current_user:"<username>"
    }
`

# Command List
 - login
    gator login <username>
    logins into a registered user

- register
    gator register <username>
    register user to the application

- reset
    gator reset
    resets the database tables

- users
    gator users
    gets all registered users

- agg
    gator agg <time>
    pull posts from a registered feeds pulling every <time> interval

- addfeed
    gator addfeed <feed name> <url>
    adds feed to list of registered feeds

- feeds
    gator feeds
    gets all registered feeds

- follow
    gator follow <url>
    follows the feed for the current user

- following
    gator following
    gets all feeds that the current user is following

- unfollow
    gator unfollow <url>
    unfollows the feed for the current user

- browse
    gator browse
    shows all registerd posts from the agg command
