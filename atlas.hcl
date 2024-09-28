env "gorm" {
  url = "mysql://root:swapnilkant11@tcp(localhost:3306)/project_database?multiStatements=true"
  
  migration {
    dir = "file://migrations"
  }
  
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
