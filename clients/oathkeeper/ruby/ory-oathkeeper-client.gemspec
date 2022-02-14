# -*- encoding: utf-8 -*-

=begin
#Ory Oathkeeper API

#Documentation for all of Ory Oathkeeper's APIs. 

The version of the OpenAPI document: v0.38.20-beta.1
Contact: hi@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 5.4.0

=end

$:.push File.expand_path("../lib", __FILE__)
require "ory-oathkeeper-client/version"

Gem::Specification.new do |s|
  s.name        = "ory-oathkeeper-client"
  s.version     = OryOathkeeperClient::VERSION
  s.platform    = Gem::Platform::RUBY
  s.authors     = ["ORY GmbH"]
  s.email       = ["opensource@ory.sh"]
  s.homepage    = "https://www.ory.sh"
  s.summary     = "Ory Oathkeeper API Ruby Gem"
  s.description = "Documentation for all of Ory Oathkeeper's APIs. "
  s.license     = "Apache-2.0"
  s.required_ruby_version = ">= 2.4"

  s.add_runtime_dependency 'typhoeus', '~> 1.0', '>= 1.0.1'

  s.add_development_dependency 'rspec', '~> 3.6', '>= 3.6.0'

  s.files         = `find *`.split("\n").uniq.sort.select { |f| !f.empty? }
  s.test_files    = `find spec/*`.split("\n")
  s.executables   = []
  s.require_paths = ["lib"]
end
