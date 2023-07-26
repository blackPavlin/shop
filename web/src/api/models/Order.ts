/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */

export type Order = {
    id: string;
    userId: string;
    status: Order.status;
};

export namespace Order {

    export enum status {
        CREATED = 'CREATED',
        ACCEPTED = 'ACCEPTED',
        CANCELED = 'CANCELED',
    }


}
