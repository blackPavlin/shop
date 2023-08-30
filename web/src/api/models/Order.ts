/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Order = {
    id: number;
    userId: number;
    status: Order.status;
};

export namespace Order {

    export enum status {
        CREATED = 'CREATED',
        ACCEPTED = 'ACCEPTED',
        CANCELED = 'CANCELED',
    }


}

