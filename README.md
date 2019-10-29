# Bulsatcom Playlist Generator

The project builds on my previous [Bulsatcom Fire TV](https://github.com/sgloutnikov/bulsatcom-firetv) project. 


### What
Produces a m3u8 playlist of Bulsatcom channels. Unfortunately only SD streams are currently available due to provider limitations.

### Why
Watch Bulsatcom in a STB-like environment with a working EPG. Have the freedom to watch on different apps and platforms.

Example of integration with Perfect Player:

<img src="https://github.com/sgloutnikov/bulsatcom-playlist-generator/raw/master/pp.jpg" alt="" width="640" height="360">

### How

* Download the compiled version from releases or build from source. The project is written entirely in Go.
* Generate your password hash using this [jar](https://github.com/sgloutnikov/bulsatcom-firetv/raw/master/bscEncrypt.jar) from
the Fire TV App project: `java -jar bscEncrypt.jar "username" "password"` (Note: quotes around username and password in case of special characters.)

**Be sure to have the Java Unlimited Strength Jurisdiction Policy Files installed**
* Edit `config.toml` accordingly and run bsc-generator.
* Use the generated bulsatcom.m3u8 playlist in your favorite player.

**Note:** The playlist links are valid for 10 hours after they are generated and the generator has to be re-ran again to grab active links. Use a scheduled task or cron to automate the process.


##### TODO
* Handle the password hash automatically.
