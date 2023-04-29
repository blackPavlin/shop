/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type User = {
    id: number;
    name: string;
    email: string;
    phone: string;
    role: User.role;
};

export namespace User {

    export enum role {
        USER = 'USER',
        ADMIN = 'ADMIN',
    }


}
