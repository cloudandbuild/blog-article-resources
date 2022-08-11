# -*- coding: utf-8 -*-

from flask import Flask, request
import papermill as pm
import scrapbook as sb
import os
import uuid

app = Flask(__name__, static_url_path="")


@app.route("/papermill/<notebook_name>/execute", methods=["POST"])
def papermill_execute(notebook_name):
    output = os.path.join('tmp', uuid.uuid4(), f'{notebook_name}.out.ipynb')
    # papermillを実行する
    pm.execute_notebook(
        f'notebooks/{notebook_name}.ipynb',
        output,
        parameters=request.json
    )

    # 結果をスクラップブックで取得する
    res = sb.read_notebook(output).scraps['result']

    return res.data


if __name__ == "__main__":
    app.run(host="0.0.0.0")
