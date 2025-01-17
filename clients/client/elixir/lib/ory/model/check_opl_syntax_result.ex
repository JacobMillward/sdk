# NOTE: This file is auto generated by OpenAPI Generator 6.2.1 (https://openapi-generator.tech).
# Do not edit this file manually.

defmodule Ory.Model.CheckOplSyntaxResult do
  @moduledoc """
  
  """

  @derive [Poison.Encoder]
  defstruct [
    :errors
  ]

  @type t :: %__MODULE__{
    :errors => [Ory.Model.ParseError.t] | nil
  }
end

defimpl Poison.Decoder, for: Ory.Model.CheckOplSyntaxResult do
  import Ory.Deserializer
  def decode(value, options) do
    value
    |> deserialize(:errors, :list, Ory.Model.ParseError, options)
  end
end

