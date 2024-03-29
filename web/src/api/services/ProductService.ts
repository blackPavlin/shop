/* generated using openapi-typescript-codegen -- do no edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CreateProductRequest } from '../models/CreateProductRequest';
import type { Product } from '../models/Product';
import type { ProductList } from '../models/ProductList';

import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';

export class ProductService {

    /**
     * Получить список товаров
     * @returns ProductList OK
     * @throws ApiError
     */
    public static getProduct(): CancelablePromise<ProductList> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/product',
            errors: {
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Создать товар
     * @param requestBody
     * @returns Product OK
     * @throws ApiError
     */
    public static postProduct(
        requestBody?: CreateProductRequest,
    ): CancelablePromise<Product> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/product',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Получить детальную карточку товара
     * @param productId ID товара
     * @returns Product OK
     * @throws ApiError
     */
    public static getProduct1(
        productId: string,
    ): CancelablePromise<Product> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/product/{productId}',
            path: {
                'productId': productId,
            },
            errors: {
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Изменить товар
     * @param productId ID товара
     * @param requestBody
     * @returns Product OK
     * @throws ApiError
     */
    public static patchProduct(
        productId: string,
        requestBody?: CreateProductRequest,
    ): CancelablePromise<Product> {
        return __request(OpenAPI, {
            method: 'PATCH',
            url: '/product/{productId}',
            path: {
                'productId': productId,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                400: `Bad Request`,
                401: `Unauthorized`,
                404: `Not Found`,
                500: `Internal Server Error`,
            },
        });
    }

    /**
     * Удалить товар
     * @param productId ID товара
     * @returns void
     * @throws ApiError
     */
    public static deleteProduct(
        productId: string,
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/product/{productId}',
            path: {
                'productId': productId,
            },
            errors: {
                401: `Unauthorized`,
                500: `Internal Server Error`,
            },
        });
    }

}
