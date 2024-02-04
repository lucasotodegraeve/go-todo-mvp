table "items" {
  schema = schema.public
  column "id" {
    null = false
    type = integer
    identity {
        generated = ALWAYS
        start = 0
        increment = 1
    }
  }
  column "completed" {
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
