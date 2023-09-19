/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { OrderDetailed } from '../models/OrderDetailed';
import type { OrderList } from '../models/OrderList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class OrderService {

    /**
     * Получить список заказов пользователя
     * @returns OrderList OK
     * @throws ApiError
     */
    public static getOrder(): CancelablePromise<OrderList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/order',
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Создать заказ
     * @returns OrderDetailed OK
     * @throws ApiError
     */
    public static postOrder(): CancelablePromise<OrderDetailed> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/order',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Получить детальную информацию о заказе
     * @param orderId ID заказа
     * @returns OrderDetailed OK
     * @throws ApiError
     */
    public static getOrder1(
        orderId: number,
    ): CancelablePromise<OrderDetailed> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/order/{orderId}',
            path: {
                'orderId': orderId,
            },
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

}
