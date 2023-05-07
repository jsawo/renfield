interface AppConfig {
  Currentproject: string,
  Projects: Record<string, Project>,
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
    ActiveTab: string,
  },
  JSONTools: {
    Tabs: Array<Tab>,
    ActiveTab: string,
    ActiveTool: string,
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