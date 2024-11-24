from flask import Flask, request, jsonify
import subprocess

app = Flask(__name__)

# フロントエンド用ルート
@app.route('/')
def index():
    return app.send_static_file('index.html')  # static/index.html を提供

# シェルスクリプトを実行するエンドポイント
@app.route('/run-script', methods=['POST'])
def run_script():
    try:
        # クライアントから送信されたパラメータを取得
        data = request.json
        param = data.get('param', '')
        if param == '':
            param = 'main'  # デフォルト値はmain
        print(f"Received parameter: {param}")  # ログに出力

        # サーバ上のスクリプトをパラメータ付きで実行
        result = subprocess.run(
            ['/home/isucon/ansible_frontend/script.sh', param],  # paramを引数として渡す
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE,
            text=True
        )

        # 実行結果を返す
        return jsonify({
            'status': 'success',
            'stdout': result.stdout,
            'stderr': result.stderr
        })
    except Exception as e:
        print(f"Error: {e}")  # エラーログを出力
        return jsonify({
            'status': 'error',
            'message': str(e)
        }), 500

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)
