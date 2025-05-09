resource "cohesity_backup_policy" "test_policy" {
  name= "testPolicy"
  backup_policy{
    regular{
      incremental{
          unit = "Days"
          day_schedule {
            frequency = 3
          }
      }
      retention{
        unit = "Weeks"
        duration = 1
      }
    }
  }
}