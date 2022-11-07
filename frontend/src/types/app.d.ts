interface AppConfig {
  Currentproject: string,
  Projects: Array<Project>,
  Tags: Array<Tag>,
}

interface Project {
  Id: string,
  Name: string,
  Path: string,
  Tag: string,
  Command: string,
  Tinker: {
    Tabs: Array<Tab>,
  }
}

interface Tab {
  Id: string,
  Name: string,
}

interface Tag {
  Label: string,
  Color: string,
}