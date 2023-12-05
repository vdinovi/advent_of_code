#!/bin/usr/env ruby

require "set"

class Entry
  class << self
    def from_string(input)
      signals, output = input.split("|").map(&:split)
      new(signals, output)
    end
  end

  attr_reader :signals, :output, :key

  def initialize(signals, output)
    @signals = signals
    @output = output
    @key = nil
  end

  def solve
    clues = @signals.reduce({}) do |hash, code|
      hash[code.length] = code.chars.to_set
      hash
    end
    @key = @output.reduce({}) do |hash, code|
      set = code.chars.to_set
      case set.length
      when 2
        hash[set] = 1
      when 3
        hash[set] = 7
      when 4
        hash[set] = 4
      when 7
        hash[set] = 8
      when 5
        if (set & clues[2]).length == 2
          hash[set] = 3
        elsif (set & clues[4]).length == 2
          hash[set] = 2
        else
          hash[set] = 5
        end
      when 6
        if (set & clues[2]).length == 1
          hash[set] = 6
        elsif (set & clues[4]).length == 4
          hash[set] = 9
        else
          hash[set] = 0
        end
      else
        raise RuntimeError, "invalid code length #{set.length} for #{code}"
      end
      hash
    end
  end

  def decode
    Integer(@output.map { |code| @key[code.chars.to_set].to_s }.join, 10)
  end

  def to_s
    [@signals.join(" "), " | ", @output.join(" ")].join
  end
end


sum = STDIN.each_line.reduce(0) do |sum, line|
  entry = Entry.from_string(line)
  entry.solve
  value = entry.decode
  puts "#{entry} => #{value}"
  sum + value
end
puts "Total: #{sum}"
