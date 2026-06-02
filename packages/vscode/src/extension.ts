import * as vscode from "vscode"
import * as cp from "child_process"

export function activate(context: vscode.ExtensionContext) {
  const output = vscode.window.createOutputChannel("Codelo")
  output.appendLine("Codelo đã được kích hoạt")

  context.subscriptions.push(
    vscode.commands.registerCommand("codelo.chat", () => {
      const root = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "."
      const t = vscode.window.createTerminal({ name: "Codelo Chat", cwd: root, env: { CODELO_LANG: "vi" } })
      t.show(); t.sendText("codelo chat")
    }),
    vscode.commands.registerCommand("codelo.run", async () => {
      const task = await vscode.window.showInputBox({ prompt: "Nhập tác vụ", placeHolder: "vd: viết hàm kiểm tra số nguyên tố" })
      if (!task) return
      const root = vscode.workspace.workspaceFolders?.[0]?.uri.fsPath || "."
      const child = cp.spawn("codelo", ["run", "--auto", task], { cwd: root, env: { ...process.env, CODELO_LANG: "vi" } })
      output.show()
      child.stdout?.on("data", (d: Buffer) => output.append(d.toString()))
      child.stderr?.on("data", (d: Buffer) => output.append(d.toString()))
      child.on("close", (code) => { output.appendLine(`Hoàn thành (code: ${code})`) })
    }),
    vscode.commands.registerCommand("codelo.setup", () => {
      const t = vscode.window.createTerminal("Codelo Setup")
      t.show(); t.sendText("codelo setup")
    })
  )
}

export function deactivate() {}