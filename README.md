# Bulsatcom Playlist Generator

The project builds on my previous [Bulsatcom Fire TV](https://github.com/sgloutnikov/bulsatcom-firetv) project. 


### What
Produces a m3u8 playlist of available Bulsatcom channels available with your account. Unfortunately only SD streams
are currently available due to provider limitations.

### Why
Watch Bulsatcom in a STB-like environment with a working EPG. Have the freedom to watch on different apps and more platforms.

Example of integration with Perfect Player:

<img src="https://raw.githubusercontent.com/sgloutnikov/bulsatcom-firetv/master/screen1.png" alt="" width="640" height="360">

### How

* Download the compiled version from releases or build from source. The project is written entirely in Go.
* Generate your password hash using this [jar](https://github.com/sgloutnikov/bulsatcom-firetv/raw/master/bscEncrypt.jar) from
the Fire TV App project. **Be sure to have the Java Unlimited Strength Jurisdiction Policy Files installed**
* Edit `config.toml` accordingly and run bsc-generator.
* Use the generated bulsatcom.m3u8 playlist in your favorite player.

**Note:** The playlist links are valid for 10 hours after they are generated and the generator has to be re-ran again 
in order to have working links. Use a scheduled task or cron to automate the process.


##### TODO
* Handle the password hash internally automatically.
