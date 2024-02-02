table "items" {
  schema = schema.public
  column "id" {
    null = false
    type = integer
    identity {
        generated = ALWAYS
        start = 10
        increment = 10
    }
  }
  column "complete" {
    type = boolean
    default = false
  }
  column "name" {
    type = character_varying(100)
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "standard public schema"
}
