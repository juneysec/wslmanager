/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */

export interface paths {
    "/distributions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * ディストリビューションのリストを取得する
         * @description ディストリビューション
         */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description ディストリビューションリスト */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": Record<string, never>[];
                    };
                };
                /** @description Bad Request */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        put?: never;
        /**
         * ディストリビューションをインポートして新規作成する
         * @description ディストリビューションをインポートして新規作成する
         */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["request_distribution_post"];
                };
            };
            responses: {
                /** @description ディストリビューションリスト */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": Record<string, never>[];
                    };
                };
                /** @description Bad Request */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Not Found */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/distributions/{distribution}": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * ディストリビューションを取得する
         * @description ディストリビューションを取得する
         */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    distribution: string;
                };
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description ディストリビューション */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_distribution"];
                    };
                };
                /** @description Not Found */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        /**
         * ディストリビューションのコマンド実行
         * @description ディストリビューションのコマンド実行
         */
        put: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    distribution: string;
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["request_distribution_put"];
                };
            };
            responses: {
                /** @description ディストリビューションリスト */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": Record<string, never>[];
                    };
                };
                /** @description Bad Request */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Not Found */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        post?: never;
        /**
         * ディストリビューションの削除
         * @description ディストリビューションの削除
         */
        delete: {
            parameters: {
                query?: never;
                header?: never;
                path: {
                    distribution: string;
                };
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["request_distribution_delete"];
                };
            };
            responses: {
                /** @description ディストリビューションリスト */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": Record<string, never>[];
                    };
                };
                /** @description Bad Request */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Not Found */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
    "/online-distributions": {
        parameters: {
            query?: never;
            header?: never;
            path?: never;
            cookie?: never;
        };
        /**
         * オンラインディストリビューションのリスト取得
         * @description オンラインディストリビューションのリスト取得
         */
        get: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: never;
            responses: {
                /** @description オンラインディストリビューションリスト */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": Record<string, never>[];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        put?: never;
        /**
         * オンラインディストリビューションをインストール
         * @description オンラインディストリビューションをインストール
         */
        post: {
            parameters: {
                query?: never;
                header?: never;
                path?: never;
                cookie?: never;
            };
            requestBody?: {
                content: {
                    "application/json": components["schemas"]["request_online_distribution_post"];
                };
            };
            responses: {
                /** @description ディストリビューションリスト */
                200: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": Record<string, never>[];
                    };
                };
                /** @description Bad Request */
                400: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Not Found */
                404: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
                /** @description Internal Server Error */
                500: {
                    headers: {
                        [name: string]: unknown;
                    };
                    content: {
                        "application/json": components["schemas"]["response_error"];
                    };
                };
            };
        };
        delete?: never;
        options?: never;
        head?: never;
        patch?: never;
        trace?: never;
    };
}
export type webhooks = Record<string, never>;
export interface components {
    schemas: {
        response_distribution: {
            /**
             * @description デフォルトの場合true
             * @example true
             */
            isDefault: boolean;
            /**
             * @description ディストリビューション名
             * @example Ubuntu
             */
            name: string;
            /**
             * @description 状態
             * @example Running
             */
            state: string;
            /**
             * @description WSLのバージョン
             * @example 2
             */
            version: string;
            /**
             * @description VHDファイルのベースパス
             * @example C:\WSL\Ubuntu2
             */
            vhdPath?: string;
            /**
             * @description VHDファイルのサイズ
             * @example 1056250135
             */
            vhdSize?: number;
        };
        response_error: {
            /**
             * @description エラーコード
             * @example 101
             */
            code?: string;
            /**
             * @description エラーメッセージ
             * @example ディストリビューションが見つかりませんでした。
             */
            message?: string;
        };
        request_distribution_post: {
            /**
             * @description ディストリビューションの名前
             * @example Ubuntu 2
             */
            name: string;
            /**
             * @description VHDファイルの作成場所
             * @example c:\wsl\Ubuntu2
             */
            vhdPath?: string;
            /**
             * @description インポート元のパス
             * @example c:\wsl\ubuntu2.tar.gz
             */
            sourcePath?: string;
        };
        request_distribution_put: {
            /**
             * @description コマンド
             * @example run
             * @enum {string}
             */
            command: "start" | "stop" | "shell" | "export" | "set-default" | "move-vhd" | "open-vhd";
            /**
             * @description コマンドが export/move-vhd の場合に、エクスポート先/移動先
             * @example c:\wsl\exp1
             */
            path?: string;
        };
        request_distribution_delete: {
            /**
             * @description ディストリビューション名
             * @example Ubuntu
             */
            name: string;
        };
        response_online_distribution: {
            /**
             * @description ディストリビューション名
             * @example Debian
             */
            name: string;
            /**
             * @description フレンドリー名
             * @example Debian GNU/Linux
             */
            friendlyName?: string;
            /**
             * @description 既にインストールされている場合は true
             * @example true
             */
            alreadyInstalled?: boolean;
        };
        request_online_distribution_post: {
            /**
             * @description ディストリビューションの名前
             * @example Ubuntu
             */
            name: string;
        };
    };
    responses: never;
    parameters: never;
    requestBodies: never;
    headers: never;
    pathItems: never;
}
export type $defs = Record<string, never>;
export type operations = Record<string, never>;
