![logo](http://buildkite.github.io/terminal/images/logo.svg)

[![Gem Version](https://badge.fury.io/rb/terminal.png)](https://rubygems.org/gems/terminal)

Terminal is a Go library (with a Ruby gem wrapper) for converting arbitrary shell output (with ANSI) into beautifully rendered HTML. See http://en.wikipedia.org/wiki/ANSI_escape_code for more information about ANSI Terminal Control Escape Sequences.

## Go Installation

Assuming a `$GOPATH/bin` that's globally accessible, run:

```bash
go install github.com/buildbox/terminal/cmd/ansi2html
```

This will give you the `ansi2html` command. It can be called directly or used by the Ruby Gem. It's called `ansi2html` and not `terminal` as installing something called `terminal` globally might confuse people looking for an actual terminal.

## Ruby Installation

Run the `go install` command above and add this line to your application's Gemfile:

```ruby
gem 'terminal'
```

And then execute:

```bash
$ bundle
```

Or install it yourself as:

```bash
gem install terminal
```

## Usage

```ruby
Terminal.render("...")
```

### Rails Integration

You can use Terminal directly within your Ruby on Rails application. First require the gem
in your Gemfile:

```ruby
gem "terminal"
```

Then in your `app/assets/application.css` file:

```css
/* require "terminal" */
```

Now in your views:

```
<div class="term-container"><%= Terminal.render(output) %></div>
```

### Emojis :+1:

The Ruby Gem wrapper of Terminal converts unicode to proper `<img>` tags. We use the [gemoji](https://github.com/github/gemoji) gem to do this. The path to the assets can be customized by passing the `:emoji_asset_path` option to `Terminal.render`

```ruby
Terminal.render(output, emoji_asset_path: "https://your.cdn.com/images/emoji")
```

### Ruby Command Line

Terminal ships with a command line utility. For example, you can pipe `rspec` output to it:

```bash
rspec --color --tty | terminal
```

Or use output saved earlier:

```bash
rspec --tty --color > output.txt
terminal output.txt
```

With `rspec`, you'll need to use the `--tty` and `--color` options to force it to output colors.

We also provide a utility to preview the rendered version in a web browser. Simply append `--preview` to the command,
and when the render has finished, it will open in your web browser with a before/after show.

```bash
rspec --color --tty | terminal --preview
```

![preview mode](http://buildkite.github.io/terminal/images/preview.png)

### With the Buildkite API

You can use the `job_url` returned by the [Builds API](https://buildkite.com/docs/api/builds) to pipe a job's log directly into terminal, for example:

```bash
JOB_LOG_URL="https://api.buildkite.com/v1/accounts/[account]/projects/[project]/builds/[build]/jobs/[job]/log.txt?api_key=[api-key]"
curl $JOB_LOG_URL | terminal --preview
```

## Generating Fixtures

To generate a fixture, first create a test case inside the `examples` folder. See the `curl.sh`
file as an example. You can then generate a `.raw` and `.rendered` file by running:

```bash
./generate curl.sh
```

You should then move the `raw` and `rendered` files to the `fixtures` folder.

```bash
mv examples/*{raw,rendered} fixtures
```

## Benchmarking

Run `go test -bench .` to see raw Go performance. The `npm` test is the focus: this best represents the kind of use cases the original code was developed against. As a guide, this test was 80ms per iteration on an 2013 Retina MBP, and was 2500 ms per iteration in the original pure Ruby implementation.

Run `script/benchmark` to see performance when pushed through the gem. Note that invocation of the binary itself applies a performance ceiling, but this is compensated for by performance on larger inputs.

## TODO

 * Have the Go version handle UTF8 enforcement
 * Have the Go version handle Emoji
 * Build the Go binary and place in gem directory as part of gem install

## Contributing

1. Fork it ( https://github.com/[my-github-username]/terminal/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request

## Licence

> Copyright (c) 2015 Keith Pitt, Tim Lucas, Buildkite Pty Ltd
>
> MIT License
>
> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the
> "Software"), to deal in the Software without restriction, including
> without limitation the rights to use, copy, modify, merge, publish,
> distribute, sublicense, and/or sell copies of the Software, and to
> permit persons to whom the Software is furnished to do so, subject to
> the following conditions:
>
> The above copyright notice and this permission notice shall be
> included in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
> MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
> NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
> LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
> OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
> WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
