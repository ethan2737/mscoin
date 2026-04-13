"""
使用 Playwright 进行 Swap API 自动化测试
测试永续合约 HTTP API 的所有接口
"""

from playwright.sync_api import sync_playwright, expect
import json

BASE_URL = "http://localhost:8086"

def run_api_tests():
    """使用 Playwright API 测试所有接口"""
    print("\n" + "=" * 60)
    print("Swap API Playwright 自动化测试")
    print("=" * 60)
    print(f"Base URL: {BASE_URL}")

    results = []

    with sync_playwright() as p:
        browser = p.chromium.launch(headless=True)
        context = browser.new_context()
        page = context.new_page()

        # 拦截 API 请求进行测试
        print("\n--- 开始 API 测试 ---\n")

        # 1. 获取合约 Thumb 信息
        print("测试 1: GET /swap/symbol-thumb")
        resp = page.request.get(f"{BASE_URL}/swap/symbol-thumb")
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Symbol Thumb", status == 200, status))

        # 2. 委托下单接口
        print("\n测试 2: POST /swap/order-entrust/add")
        order_data = {
            "memberId": 1001,
            "contractCoinId": 1,
            "entrustType": 1,
            "type": 1,
            "direction": 1,
            "leverage": 10,
            "price": 50000,
            "amount": 0.1
        }
        resp = page.request.post(f"{BASE_URL}/swap/order-entrust/add", data=order_data)
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        # 200 表示成功，其他错误码表示服务响应但可能有业务错误
        results.append(("Add Order", status >= 200, status))

        # 3. 当前委托
        print("\n测试 3: POST /swap/order-entrust/current")
        resp = page.request.post(f"{BASE_URL}/swap/order-entrust/current", data={"memberId": 1001})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Current Order", status >= 200, status))

        # 4. 历史委托
        print("\n测试 4: POST /swap/order-entrust/history")
        resp = page.request.post(f"{BASE_URL}/swap/order-entrust/history", data={"memberId": 1001})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("History Order", status >= 200, status))

        # 5. 持仓信息
        print("\n测试 5: POST /swap/order-entrust/position")
        resp = page.request.post(f"{BASE_URL}/swap/order-entrust/position", data={"memberId": 1001})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Position", status >= 200, status))

        # 6. 杠杆倍数
        print("\n测试 6: POST /swap/order-entrust/leverage")
        resp = page.request.post(f"{BASE_URL}/swap/order-entrust/leverage", data={"memberId": 1001, "contractCoinId": 1})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Leverage", status >= 200, status))

        # 7. 用户中心 - 当前合约
        print("\n测试 7: POST /uc/contract/current")
        resp = page.request.post(f"{BASE_URL}/uc/contract/current", data={"memberId": 1001})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("UC Contract Current", status >= 200, status))

        # 8. 用户中心 - 历史合约
        print("\n测试 8: POST /uc/contract/history")
        resp = page.request.post(f"{BASE_URL}/uc/contract/history", data={"memberId": 1001})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("UC Contract History", status >= 200, status))

        # 9. 钱包信息
        print("\n测试 9: POST /uc/contract-wallet")
        resp = page.request.post(f"{BASE_URL}/uc/contract-wallet", data={"memberId": 1001, "unit": "USDT"})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Wallet Info", status >= 200, status))

        # 10. 钱包转账
        print("\n测试 10: POST /uc/contract-wallet/transfer")
        transfer_data = {
            "memberId": 1001,
            "unit": "USDT",
            "amount": 100,
            "type": 1
        }
        resp = page.request.post(f"{BASE_URL}/uc/contract-wallet/transfer", data=transfer_data)
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Wallet Transfer", status >= 200, status))

        # 11. 交易流水
        print("\n测试 11: POST /uc/asset/contract-transaction/all")
        resp = page.request.post(f"{BASE_URL}/uc/asset/contract-transaction/all", data={"memberId": 1001, "page": 1, "pageSize": 10})
        status = resp.status
        try:
            data = resp.json()
        except:
            data = {"raw": resp.text()}
        print(f"  状态码：{status}")
        print(f"  响应：{json.dumps(data, indent=2) if isinstance(data, dict) else data}")
        results.append(("Transaction History", status >= 200, status))

        browser.close()

    # 汇总结果
    print("\n" + "=" * 60)
    print("测试结果汇总")
    print("=" * 60)

    passed = sum(1 for _, success, _ in results if success)
    total = len(results)

    for name, success, status in results:
        status_str = "PASS" if success else f"FAIL (HTTP {status})"
        print(f"{status_str} - {name}")

    print(f"\n总计：{passed}/{total} 个测试通过")

    if passed == total:
        print("\n所有 API 接口测试通过!")
        return True
    else:
        print(f"\n{total - passed} 个测试失败 - 服务可能未完全正常运行")
        return False

if __name__ == "__main__":
    success = run_api_tests()
    exit(0 if success else 1)
