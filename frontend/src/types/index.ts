export interface TaskView {
    id: string;
    description: string;
    checked: boolean;
    children: TaskView[];
    action:string,
    root: boolean;
}
export type Task  = TaskView

interface Attrs {
    [key: string]: string;
}
export enum ItemType {
    Doc = "doc",
    TaskList = "taskList",
    TaskItem = "taskItem",
    Paragraph = "paragraph",
    Text = "text",
}
export type Item = DocJson;
export interface DocJson {
    type: ItemType;
    content: Array<Item>;
    attrs?:Attrs;
    text?: string;
}