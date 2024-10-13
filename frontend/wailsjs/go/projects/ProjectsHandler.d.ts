// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {projects} from '../models';

export function CreateNotes(arg1:number,arg2:string):Promise<void>;

export function CreateProject(arg1:projects.ProjectReq):Promise<void>;

export function DeleteProject(arg1:number):Promise<void>;

export function EditPersonalNotes(arg1:number,arg2:string):Promise<void>;

export function EditProject(arg1:projects.Project):Promise<void>;

export function GetAllProjects():Promise<Array<projects.Project>>;

export function GetChildProjectsByParentID(arg1:number):Promise<Array<projects.Project>>;

export function GetNotesByProjectID(arg1:number):Promise<string>;

export function GetProjectByID(arg1:string):Promise<projects.ProjectRes>;

export function GetProjectsByStatus(arg1:number):Promise<Array<projects.ProjectRes>>;

export function GetProjectsByUrgency(arg1:number):Promise<Array<projects.ProjectRes>>;

export function SearchProjectAssignee(arg1:projects.SearchReq):Promise<Array<projects.User>>;

export function SearchProjects(arg1:projects.SearchReq):Promise<Array<projects.Project>>;
