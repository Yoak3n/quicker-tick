import type { Action } from '@/types';


export interface  RowData {
    key:string;
    title: string;
    description: string;
    tags: string[];
    aciton: Action;
    date: string;
}


export function StartTask(row:RowData) {
    console.log(row);
    
}