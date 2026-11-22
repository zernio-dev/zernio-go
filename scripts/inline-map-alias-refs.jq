.components.schemas as $schemas
| (.. | objects | select(.anyOf? | type == "array") | .anyOf[]) |= (
    (.["$ref"]? // "") as $ref
    | if ($ref | startswith("#/components/schemas/")) then
        $schemas[$ref | ltrimstr("#/components/schemas/")] as $target
        | if ($target.type == "object" and ($target | has("additionalProperties")) and ($target | has("properties") | not))
          then $target else . end
      else . end
  )
