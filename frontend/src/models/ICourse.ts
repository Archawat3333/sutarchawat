import { QualificationsInterface } from "./IQualification";
import { MajorsInterface } from "./IMajor";

export interface CoursesInterface {

    Course_ID : string,
   
    Course_Name: string;
   
    Datetime: Date | null;

    Qualification_ID: string;
    Qualification: QualificationsInterface

    Major_ID: string;
    Major: MajorsInterface
    

   }