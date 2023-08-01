/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Cart } from '../models/Cart';
import type { CartProduct } from '../models/CartProduct';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class CartService {

    /**
     * Получить корзину пользователя
     * @returns Cart OK
     * @throws ApiError
     */
    public static getCart(): CancelablePromise<Cart> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/cart',
        });
    }

    /**
     * Добавить товар в корзину
     * @param requestBody
     * @returns Cart OK
     * @throws ApiError
     */
    public static postCart(
        requestBody?: CartProduct,
    ): CancelablePromise<Cart> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/cart',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Изменить количество товара в корзине
     * @param requestBody
     * @returns Cart OK
     * @throws ApiError
     */
    public static patchCart(
        requestBody?: CartProduct,
    ): CancelablePromise<Cart> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/cart',
            body: requestBody,
            mediaType: 'application/json',
        });
    }

    /**
     * Очистить корзину
     * @returns Cart OK
     * @throws ApiError
     */
    public static deleteCart(): CancelablePromise<Cart> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/cart',
        });
    }

}
