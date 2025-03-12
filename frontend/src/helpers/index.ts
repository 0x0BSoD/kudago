function formatBytes(bytes: number): [number, string] {
    const unit = 1024;
    if (bytes < unit) {
        return [bytes, 'B'];
    }

    const units = ['Kb', 'Mb', 'Gb', 'Tb', 'Pb', 'Eb'];
    let div = unit;
    let exp = 0;

    for (let n = bytes / unit; n >= unit; n /= unit) {
        div *= unit;
        exp++;
    }

    const value = parseFloat((bytes / div).toFixed(2));
    return [value, units[exp]];
}

export { formatBytes }