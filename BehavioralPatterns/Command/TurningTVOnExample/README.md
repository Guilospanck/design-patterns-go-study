# Turning TV and Radio Example
This is an example of the `Command` pattern.
A person will turn a device (tv or radio) on or off using a button (that can be a remote control or physically touching the device).

## Commands
 - `ICommand` defines the common interface for all commands.
 - `OnCommand` implements the `ICommand` interface for turning devices on.
 - `OffCommand` implements the `ICommand` interface for turning devices off.

## Receivers
The receivers are the `devices`. A receiver is someone that will receive the action from the `command`.
  - `IDevice` defines the common interface for all devices.
  - `TV` implements the `IDevice` interface for TV devices.
  - `Radio` implements the `IDevice` interface for Radio devices.

## Sender
A sender is something responsible for triggering the action (click, send, copy ...).
This is when the `Command` comes into play, converting that operation into an object.
  - `Button` is our sender.