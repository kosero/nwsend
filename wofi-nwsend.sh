#!/usr/bin/env bash
message=$(wofi --dmenu --prompt "Message...")
nwsend "$message"
