export namespace main {
	
	export class User {
	    id: number;
	    username: string;
	    email: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.email = source["email"];
	    }
	}

}

export namespace projects {
	
	export class Project {
	    id: number;
	    parentID: number;
	    title: string;
	    description: string;
	    status: number;
	    assigneeID: number;
	    urgency: number;
	    notes: string;
	    startDate: string;
	    endDate: string;
	    // Go type: time
	    createdAt: any;
	    // Go type: time
	    updatedAt: any;
	
	    static createFrom(source: any = {}) {
	        return new Project(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.parentID = source["parentID"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	        this.assigneeID = source["assigneeID"];
	        this.urgency = source["urgency"];
	        this.notes = source["notes"];
	        this.startDate = source["startDate"];
	        this.endDate = source["endDate"];
	        this.createdAt = this.convertValues(source["createdAt"], null);
	        this.updatedAt = this.convertValues(source["updatedAt"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ProjectReq {
	    parentID: number;
	    title: string;
	    description: string;
	    status: number;
	    assigneeID: number;
	    urgency: number;
	    notes: string;
	    startDate: string;
	    endDate: string;
	
	    static createFrom(source: any = {}) {
	        return new ProjectReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.parentID = source["parentID"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.status = source["status"];
	        this.assigneeID = source["assigneeID"];
	        this.urgency = source["urgency"];
	        this.notes = source["notes"];
	        this.startDate = source["startDate"];
	        this.endDate = source["endDate"];
	    }
	}
	export class User {
	    id: number;
	    username: string;
	    email: string;
	
	    static createFrom(source: any = {}) {
	        return new User(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.username = source["username"];
	        this.email = source["email"];
	    }
	}
	export class ProjectRes {
	    project?: Project;
	    user?: User;
	
	    static createFrom(source: any = {}) {
	        return new ProjectRes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.project = this.convertValues(source["project"], Project);
	        this.user = this.convertValues(source["user"], User);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SearchReq {
	    text: string;
	
	    static createFrom(source: any = {}) {
	        return new SearchReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.text = source["text"];
	    }
	}

}

