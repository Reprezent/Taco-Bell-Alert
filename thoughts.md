### Thoughts
 - Need a way to store the last run time/data on disk
        Would take the format of like
            -- Time
            -- String of items
            -- (Price)?
 - Store multiple?
 - Simple thing could be a .tbell file and then, just store them with the timestamp as the filename
 - Would need rotation/limits, as to not fill up stuff (But I doubt that)

#### Outputs
 - Make an outputs interface
 - Gmail email message (Free)
    Just use smtp stuff, gonna be annoying
 - Use an SMS (Twillo/Bandwidth)

#### Inputs
 - Right now we have Tbell, but what if we want more?
 - Make a scrape interface to check for new stuff
    Would need a name, and then a "scrape" funcs
    Name is needed so we can sort them in the "database"
    scrape is to just scrape the sites

