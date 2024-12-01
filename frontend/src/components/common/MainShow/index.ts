import type { Action } from '@/types';


export interface  RowData {
    key:string;
    checked:boolean;
    title: string;
    description: string;
    tags: string[];
    actions: Action[];
    date: string;
}


export function StartTask(row:RowData) {
    console.log(row);
}